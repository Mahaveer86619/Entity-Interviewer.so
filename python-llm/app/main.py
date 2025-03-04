from fastapi import FastAPI

app = FastAPI()

@app.post("/generate-questions/")
def generate_questions(job_description: str):
    # Implement LLM-based question generation logic
    return {"questions": ["Question 1", "Question 2"]}

@app.get("/health")
async def health_check():
    return {"status": "OK from LLM Microservice"}