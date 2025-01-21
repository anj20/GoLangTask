### README: Video Player Application with Interactive Ad Overlays

---

## **Objective**
This project aims to build a video player application with interactive ad overlays that dynamically update every 10 seconds. Users can click on the ads, which opens the respective target URL and logs metadata for analytics.

---

## **Features**

### **Frontend (React)**
1. Video player interface with overlaying ad icons/images.
2. Dynamic ad changes every 10 seconds.
3. Ads are displayed at random positions over the video.
4. Clicking an ad:
   - Opens the ad’s target URL in a new tab.
   - Sends metadata (ad ID, timestamp, user IP, and video playback time) to the backend.

### **Backend (GoLang)**
1. **APIs**:
   - `GET /ads`: Fetches a list of ads with metadata.
   - `POST /ads/click`: Logs ad click metadata.
2. **Data Logging**:
   - Ad ID.
   - Timestamp.
   - User’s IP address.
   - Video playback time.
3. Proper error handling and logging.
4. Uses SQLite as the database for simplicity and portability.

---

## **System Architecture**

### **Frontend**
- Built with React.
- Video player with ad overlays using HTML5 `<video>` tag.
- Fetches ad data from the backend using `GET /ads`.
- Logs ad click metadata with `POST /ads/click`.

### **Backend**
- Built with GoLang.
- SQLite database to store:
  - Ads metadata.
  - Logged ad click details.
- RESTful APIs for frontend communication.
- Middleware for IP logging.

### **Dockerized Environment**
- Separate Docker containers for frontend and backend.
- `docker-compose.yml` for orchestration.

---

## **Setup Instructions**

### **Prerequisites**
1. [Node.js](https://nodejs.org/) (for frontend)
2. [GoLang](https://go.dev/) (for backend)
3. [Docker](https://www.docker.com/)
4. [SQLite](https://sqlite.org/)

---

### **Running Locally**

#### **Frontend**
1. Navigate to the `frontend` directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm start
   ```
4. Access the app at `http://localhost:3000`.

#### **Backend**
1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```
2. Initialize the SQLite database:
   ```bash
   go run init_db.go
   ```
3. Start the backend server:
   ```bash
   go run main.go
   ```
4. The server runs on `http://localhost:8080`.

---

### **Running with Docker**

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```
2. Build and run the containers:
   ```bash
   docker-compose up --build
   ```
3. Access the application:
   - Frontend: `http://localhost:3000`
   - Backend API: `http://localhost:8080`

---

## **Project Structure**

### **Frontend**
```
/frontend
├── src
│   ├── components
│   ├── services
│   ├── App.js
│   ├── index.js
├── public
│   ├── index.html
├── package.json
```

### **Backend**
```
/backend
├── main.go
├── handlers
├── models
├── database
│   ├── init_db.go
├── Dockerfile
```

---

## **API Endpoints**

### **GET /ads**
Fetches a list of ads.
- **Response Example**:
  ```json
  [
    {
      "id": 1,
      "image_url": "https://example.com/image.png",
      "target_url": "https://example.com",
      "video_time": "00:10"
    }
  ]
  ```

### **POST /ads/click**
Logs ad click metadata.
- **Request Body**:
  ```json
  {
    "ad_id": 1,
    "ip_address": "192.168.1.1",
    "video_time": "00:10"
  }
  ```
- **Response**:
  ```json
  {
    "message": "Ad click logged successfully"
  }
  ```

---

## **Git Workflow**
1. **Branches**:
   - `main`: Production-ready code.
   - `feature/video-player`: Frontend development.
   - `feature/ads-api`: Backend API development.
2. **Commit Messages**:
   - Follow meaningful messages:
     ```bash
     git commit -m "Add dynamic ad overlay functionality"
     ```

---

## **Production Readiness**
1. Dockerized containers for both frontend and backend.
2. Environment variables for configuration:
   - API Base URL.
   - Database connection settings.
3. Clear instructions for local and production setup.

---

## **Evaluation Criteria**

1. **Functionality**:
   - Ads display dynamically.
   - Ad clicks log metadata correctly.
2. **Code Quality**:
   - Modular and maintainable.
   - Proper comments and best practices.
3. **Production Readiness**:
   - Dockerized setup with clear README.
4. **Bonus Features**:
   - Advanced UI/UX enhancements.
   - Analytics dashboard (optional).

---

## **Future Enhancements**
1. Add unit tests for the backend APIs.
2. Implement user sessions for personalized analytics.
3. Build an analytics dashboard to visualize ad performance.

---

