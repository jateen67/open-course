import asyncio
import httpx
from bs4 import BeautifulSoup


async def fetch_httpx():
    urls = [
        "https://vsb.mcgill.ca/vsb/getclassdata.jsp?term=202409&course_0_0=COMP-204&rq_0_0=null&t=121&e=45&nouser=1&_=1720387283792",
    ]

    async with httpx.AsyncClient() as httpx_client:
        req = [httpx_client.get(addr) for addr in urls]
        res = await asyncio.gather(*req, return_exceptions=True)

        for response in res:
            if isinstance(response, httpx.Response):
                if response.status_code == 200:
                    soup = BeautifulSoup(response.text, "xml")

                    errors_tags = soup.find_all("errors")
                    course_tags = soup.find_all("course")
                    block_tags = soup.find_all("block")

                    for course_tag in course_tags:
                        course_number = course_tag.get("key")
                        course_title = course_tag.get("title")

                    for block_tag in block_tags:
                        section = block_tag.get("disp")
                        open_seats = block_tag.get("os")
                        waitlist_taken = block_tag.get("ws")
                        waitlist_capacity = block_tag.get("wc")
                        print(f"Class: {course_number} - {course_title}")
                        print(f"Section: {section}")
                        print(f"Open seats: {open_seats}")
                        print(f"Waitlist: {waitlist_taken}/{waitlist_capacity}")

                else:
                    print(f"Error {response.status_code}: {response.text}")
            else:
                print(f"Request failed with exception: {response}")


asyncio.run(fetch_httpx())
