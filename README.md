# 🧠 Sentiment Analyzer CLI

A simple proof-of-concept project demonstrating basic AI/ML integration using a **Go CLI** and a **Python FastAPI** backend. This tool analyzes the sentiment of a given text string (Positive, Negative, or Neutral) using TextBlob.

---

## 📦 Tech Stack

- **Go** – CLI app that sends input text to the API and shows the sentiment result
- **Python** – FastAPI backend using TextBlob for sentiment analysis
- **REST API** – Communication between CLI and backend

---

## 📁 Project Structure

sentiment-analyzer-cli/ ├── cli/ # Go CLI client │ ├── main.go │ └── go.mod ├── api/ # Python FastAPI backend │ ├── main.py │ └── requirements.txt ├── .gitignore └── README.md


---

## ⚙️ Setup Instructions

### 🔧 Backend (Python + FastAPI)

1. Create and activate a virtual environment:

```bash
cd api
python3 -m venv venv
source venv/bin/activate

    Install dependencies:

pip install -r requirements.txt

    Run the API:

uvicorn main:app --reload

    The API runs on http://127.0.0.1:8000

💻 CLI (Go)

    Build and run the CLI:

cd cli
go run main.go "I love building CLI tools!"

    Output:

Sentiment: Positive

    The CLI sends the input string to the FastAPI server and displays the result.

🧪 Example

go run main.go "This app is amazing and works great!"

Returns:

Sentiment: Positive

🔮 How It Works

    CLI makes a POST request to http://localhost:8000/analyze

    FastAPI receives the text and uses TextBlob to analyze sentiment

    Sentiment polarity is mapped to:

        > 0.1 → Positive

        < -0.1 → Negative

        Otherwise → Neutral