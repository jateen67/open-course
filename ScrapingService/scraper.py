import time
import asyncio
import httpx


async def fetch_httpx():
    urls = [
        "https://en.wikipedia.org/wiki/Badlands",
        "https://en.wikipedia.org/wiki/Canyon",
        "https://en.wikipedia.org/wiki/Cave",
        "https://en.wikipedia.org/wiki/Cliff",
        "https://en.wikipedia.org/wiki/Coast",
        "https://en.wikipedia.org/wiki/Continent",
        "https://en.wikipedia.org/wiki/Coral_reef",
        "https://en.wikipedia.org/wiki/Desert",
        "https://en.wikipedia.org/wiki/Forest",
        "https://en.wikipedia.org/wiki/Geyser",
        "https://en.wikipedia.org/wiki/Mountain_range",
        "https://en.wikipedia.org/wiki/Peninsula",
        "https://en.wikipedia.org/wiki/Ridge",
        "https://en.wikipedia.org/wiki/Savanna",
        "https://en.wikipedia.org/wiki/Shoal",
        "https://en.wikipedia.org/wiki/Steppe",
        "https://en.wikipedia.org/wiki/Tundra",
        "https://en.wikipedia.org/wiki/Valley",
        "https://en.wikipedia.org/wiki/Volcano",
        "https://en.wikipedia.org/wiki/Artificial_island",
        "https://en.wikipedia.org/wiki/Lake",
    ]

    async with httpx.AsyncClient() as httpx_client:
        req = [httpx_client.get(addr) for addr in urls]

        result = await asyncio.gather(*req)


start = time.time()
asyncio.run(fetch_httpx())
end = time.time()

print("Total Consumed Time using HTTPX", end - start)
