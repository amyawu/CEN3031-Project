from fastapi import FastAPI
from fastapi.testclient import TestClient
from sample_network import Model
from pydantic import BaseModel
import uvicorn

model = Model()
app = FastAPI()


class Image(BaseModel):
    img_url: str


@app.get("/")
async def read_main():
    return {"msg": "Hello World"}


client = TestClient(app)


def test_read_main():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"msg": "Hello World"}

@app.post("/python/")
async def classify(image : Image):
    result = model.predict(image.img_url)
    return {"message": result}



if __name__ == "__main__": 
    print("MANUEL DEBUG MODE")
    uvicorn.run("main:app", host="localhost", port=8080, reload = True)
