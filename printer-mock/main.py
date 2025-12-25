import asyncio
import json
import random
import websockets
import os

GATEWAY_URL = os.getenv("GATEWAY_URL", "ws://localhost:8000/ws")
PRINTER_ID = os.getenv("PRINTER_ID", "virtual-printer-1")

async def run():
    async with websockets.connect(GATEWAY_URL) as ws:
        await ws.send(json.dumps({
            "type": "register_printer",
            "payload": {
                "printer_id": PRINTER_ID
            }
        }))
        
        print(f"[printer-mock] registered as {PRINTER_ID}")

        while True:
            telemetry = {
                "temp_nozzle": round(random.uniform(120, 220), 1),
                "temp_bed": round(random.uniform(40, 70), 1),
                "state": "printing"
            }

            await ws.send(json.dumps({
                "type": "telemetry",
                "payload": telemetry
            }))

            await asyncio.sleep(1)

asyncio.run(run())