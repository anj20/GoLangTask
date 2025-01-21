import React, { useState, useEffect } from "react";
import axios from "axios";
import { API_BASE_URL, AD_FETCH_INTERVAL, AD_POSITIONS } from "../constants";

const VideoPlayer = () => {
  const [ads, setAds] = useState([]);
  const [currentAd, setCurrentAd] = useState(null);
  const [adPosition, setAdPosition] = useState({});

  // Fetch ads from the API
  useEffect(() => {
    const fetchAds = async () => {
      try {
        const response = await axios.get(`${API_BASE_URL}/ads`);
        console.log(API_BASE_URL);
        console.log(response.data);
        setAds(response.data);
      } catch (error) {
        console.error("Error fetching ads:", error);
      }
    };
    fetchAds();
  }, []);

  // Update ad position and display a random ad every interval
  useEffect(() => {
    const interval = setInterval(() => {
      if (ads.length > 0) {
        const randomAd = ads[Math.floor(Math.random() * ads.length)];
        const randomPosition = AD_POSITIONS[Math.floor(Math.random() * AD_POSITIONS.length)];
        setCurrentAd(randomAd);
        setAdPosition(randomPosition);
      }
    }, AD_FETCH_INTERVAL);

    return () => clearInterval(interval);
  }, [ads]);

  const handleAdClick = async (ad) => {
    try {
      await axios.post(`${API_BASE_URL}/ads/click`, { ad_id: ad.id });
      window.open(ad.target_url, "_blank");
    } catch (error) {
      console.error("Error logging ad click:", error);
    }
  };

  return (
    <div className="relative w-full max-w-screen-lg mx-auto">
      <video
        className="w-full bg-black"
        controls
        src="path-to-your-video.mp4"
        autoPlay
        loop
      ></video>
      {currentAd && (
        <img
          src={currentAd.image_url}
          alt="Ad"
          className="absolute"
          style={{ ...adPosition }}
          onClick={() => handleAdClick(currentAd)}
        />
      )}
    </div>
  );
};

export default VideoPlayer;
