from fastapi import FastAPI
import os

app = FastAPI()
api_host = os.getenv("API_HOST", "http://localhost:8081")

@app.get("/")
def read_root():
    return {"message": "LLM Microservice"}

@app.post("/generate-questions/")
def generate_questions(job_description: str):
    # Implement LLM-based question generation logic
    return {"questions": ["Question 1", "Question 2"]}