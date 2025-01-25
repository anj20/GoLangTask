import React from "react";
import VideoPlayer from "./components/VideoPlayer";

const App = () => {
  return (
    <div className="bg-gray-300 min-h-screen flex flex-col items-center py-10">
      <h1
        className="text-8xl font-extrabold mb-6 text-transparent bg-clip-text bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 hover:scale-110 hover:rotate-3 transition-transform duration-300"
        title="Dynamic Video Ads"
      >
        Dynamic Video Ads
      </h1>
      <VideoPlayer />
    </div>
  );
};

export default App;
