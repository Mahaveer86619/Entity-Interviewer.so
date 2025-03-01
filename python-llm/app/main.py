from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def read_root():
    return {"message": "LLM Microservice"}

@app.post("/generate-questions/")
def generate_questions(job_description: str):
    # Implement LLM-based question generation logic
    return {"questions": ["Question 1", "Question 2"]}