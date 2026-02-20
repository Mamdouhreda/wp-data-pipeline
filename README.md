📝 WordPress Content Enhancer
WordPress Content Enhancer is a high-performance Go pipeline designed to bridge the gap between raw CMS data and AI-optimized content. It fetches posts via the WordPress REST API, enhances them using the Llama 3.3 model (via Groq), and archives them in MongoDB for downstream use in apps, newsletters, or analytics.

✨ Key Features
⚡ High-Performance Go Backend: Leveraging Go concurrency for efficient API handling.

🤖 AI-Powered Refinement: Automatically rewrites titles and summarizes content for better engagement.

🗄️ Flexible Storage: Stores enriched JSON documents in MongoDB, ready for any modern frontend.

🛡️ Robust Error Handling: Graceful fallbacks to original content if AI services are unreachable.

🧩 Modular Architecture: Easily swap out Groq for OpenAI or MongoDB for PostgreSQL.

🛠 Tech Stack
Component	Technology
Language	Go (Golang)
AI Engine	Llama 3.3 (via Groq API)
Database	MongoDB
Config	Dotenv (.env)
🚀 Getting Started
1. Prerequisites

Go 1.21 or higher

A running MongoDB instance (Local or Atlas)

A Groq API Key

2. Installation

Bash
# Clone the repository
git clone <repo-url>
cd wordpress-content-enhancer

# Install dependencies
go mod tidy

3. Configuration

Create two environment files in the root directory:

.env (AI Configuration)

Code snippet
GROQ_API_KEY=your_lp_api_key_here
WP_API_URL=https://your-site.com/wp-json/wp/v2/posts

.env.mongo (Database Configuration)

Code snippet
MONGO_URI=mongodb://admin:password@localhost:27017
MONGO_DB_NAME=wp_enhancer

4. Running the App
Bash
go run main.go
🏗 Project Structure
Plaintext
.
├── agent/            # AI logic & Groq API integration
├── db/               # MongoDB connection & CRUD operations
├── models/           # Shared Go structs for WP & DB documents
├── .env              # AI & Site secrets (git ignored)
├── .env.mongo        # DB secrets (git ignored)
├── main.go           # Orchestrator: Fetch -> Enhance -> Store
└── go.mod            # Dependencies

🔄 Workflow Logic
Ingestion: The app queries the WordPress REST API for the latest posts.

Processing: Each post is passed to the agent package. The AI analyzes the context and generates a punchier, SEO-friendly title.

Transformation: The system merges the original metadata with the AI-generated enhancements.

Persistence: The final "Enhanced Document" is indexed and saved to MongoDB.


🤝 Contributing
Contributions make the open-source community an amazing place to learn and create.

Fork the Project

Create your Feature Branch (git checkout -b feature/AmazingFeature)

Commit your Changes (git commit -m 'Add some AmazingFeature')

Push to the Branch (git push origin feature/AmazingFeature)

Open a Pull Request

