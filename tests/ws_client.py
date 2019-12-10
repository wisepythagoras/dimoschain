#!/usr/bin/env python3

import asyncio
import websockets


async def hello():
    async with websockets.connect('ws://localhost:31413') as websocket:
        msg = 'HELLO'

        await websocket.send(msg)
        print(f"> {msg}")

        greeting = await websocket.recv()
        print(f"< {greeting}")


asyncio.get_event_loop().run_until_complete(hello())
