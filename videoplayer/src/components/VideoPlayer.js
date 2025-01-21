import React, { useEffect, useState } from "react";
import AdIcon from "./AdIcon";
import "./VideoPlayer.css";

const VIDEO_URL = "https://www.w3schools.com/html/mov_bbb.mp4"; // Example video URL
const API_GET_ADS = "http://localhost:8080/ads"; // Backend API for fetching ads
const API_POST_CLICK = "http://localhost:8080/ads/click"; // Backend API for logging ad clicks

const VideoPlayer = () => {
  const [ads, setAds] = useState([]);
  const [currentAd, setCurrentAd] = useState(null);
  const [adPosition, setAdPosition] = useState({ top: "10%", left: "10%" });

  useEffect(() => {
    // Fetch ads from the API
    const fetchAds = async () => {
      try {
        const response = await fetch(API_GET_ADS);
        const data = await response.json();
        setAds(data);
      } catch (error) {
        console.error("Failed to fetch ads:", error);
      }
    };

    fetchAds();
  }, []);

  useEffect(() => {
    // Show a new ad every 10 seconds
    const interval = setInterval(() => {
      if (ads.length > 0) {
        const randomAd = ads[Math.floor(Math.random() * ads.length)];
        setCurrentAd(randomAd);
        setAdPosition({
          top: `${Math.random() * 80}%`,
          left: `${Math.random() * 80}%`,
        });
      }
    }, 10000);

    return () => clearInterval(interval);
  }, [ads]);

  const handleAdClick = async (ad) => {
    try {
      const videoElement = document.getElementById("video");
      const videoTime = videoElement.currentTime;

      await fetch(API_POST_CLICK, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          ad_id: ad.id,
          video_time: videoTime,
        }),
      });
      window.open(ad.target_url, "_blank");
    } catch (error) {
      console.error("Failed to log ad click:", error);
    }
  };

  return (
    <div className="video-container">
      <video id="video" src={VIDEO_URL} controls autoPlay />
      {currentAd && (
        <AdIcon
          ad={currentAd}
          position={adPosition}
          onClick={() => handleAdClick(currentAd)}
        />
      )}
    </div>
  );
};

export default VideoPlayer;
