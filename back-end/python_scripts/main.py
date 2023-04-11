from fastapi import FastAPI
from sample_network import Model
from pydantic import BaseModel
import uvicorn

model = Model()
app = FastAPI()


class Image(BaseModel):
    img_url: str



@app.post("/python/")
async def classify(image : Image):
    model.predict(image.img_url)
    return {"message": "Hello World"}



if __name__ == "__main__": 
    print("MANUEL DEBUG MODE")
    uvicorn.run("main:app", host="localhost", port=8080, reload = True)
