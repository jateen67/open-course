import asyncio
import json
import httpx
from bs4 import BeautifulSoup


async def fetch_httpx():
    urls = [
        "https://vsb.mcgill.ca/vsb/getclassdata.jsp?term=202409&course_0_0=COMP-307&rq_0_0=null&course_1_0=MATH-323&rq_1_0=null&t=176&e=15&nouser=1&_=1720390617241"
    ]
    courses_list = []

    async with httpx.AsyncClient() as httpx_client:
        req = [httpx_client.get(addr) for addr in urls]
        res = await asyncio.gather(*req, return_exceptions=True)

        for response in res:
            if isinstance(response, httpx.Response):
                if response.status_code == 200:
                    soup = BeautifulSoup(response.text, "xml")

                    errors_tags = soup.find_all("errors")
                    for errors_tag in errors_tags:
                        error_message = errors_tag.find("error")
                        if error_message is not None:
                            print(f"Error: {error_message.content}")
                            return

                    course_tags = soup.find_all("course")

                    for course_tag in course_tags:
                        course_number = course_tag.get("key")
                        course_title = course_tag.get("title")
                        block_tags = course_tag.findChildren("block")

                        for block_tag in block_tags:
                            section = block_tag.get("disp")
                            open_seats = block_tag.get("os")
                            waitlist_available = block_tag.get("ws")
                            waitlist_capacity = block_tag.get("wc")
                            course_object = {
                                "course_code": course_number,
                                "course_title": course_title,
                                "section": section,
                                "open_seats": open_seats,
                                "waitlist_available": int(waitlist_available),
                                "waitlist_capacity": int(waitlist_capacity),
                            }
                            courses_list.append(course_object)

                    with open("courses.json", "w") as f:
                        json.dump(courses_list, f, indent=4)

                else:
                    print(f"Error {response.status_code}: {response.text}")
                    return
            else:
                print(f"Request failed with exception: {response}")


asyncio.run(fetch_httpx())
