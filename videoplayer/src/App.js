import React from "react";
import VideoPlayer from "./components/VideoPlayer";

const App = () => {
  return (
    <div className="bg-gray-100 min-h-screen flex flex-col items-center py-10">
      <h1 className="text-3xl font-bold mb-6 text-gray-800">Dynamic Video Ads</h1>
      <VideoPlayer />
    </div>
  );
};

export default App;
