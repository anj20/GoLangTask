import React from "react";
import VideoPlayer from "./components/VideoPlayer";
import "./App.css";

const App = () => {
  return (
    <div className="app">
      <h1>Dynamic Video Ads</h1>
      <VideoPlayer />
    </div>
  );
};

export default App;
