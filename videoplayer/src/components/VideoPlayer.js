import React, { useState, useEffect, useRef } from "react";
import axios from "axios";
import { API_BASE_URL, AD_FETCH_INTERVAL, AD_POSITIONS } from "../constants";

const VideoPlayer = () => {
  const [ads, setAds] = useState([]);
  const [currentAd, setCurrentAd] = useState(null);
  const [adPosition, setAdPosition] = useState({});
  const videoRef = useRef(null);
  const [timestamp, setTimestamp] = useState(0);

  const handleTimeUpdate = () => {
    if (videoRef.current) {
      setTimestamp(videoRef.current.currentTime.toFixed(2)); // Format to 2 decimal places
    }
  };
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
      // Fetch the user's IP address
      const ipResponse = await axios.get("https://api.ipify.org?format=json");
      const { ip } = ipResponse.data;
      console.log("The ip address is:",ip);
      console.log("The curretn time is:",timestamp);
      // Send the ad click data to the backend
      await axios.post(`${API_BASE_URL}/ads/click`, {
        ad_id: ad.id,
        ip_address: ip, // Include the IP address in the payload
        timestamp:timestamp.toString()
      });
      window.open(ad.target_url, "_blank");
    } catch (error) {
      console.error("Error logging ad click:", error);
    }
  };

  return (
    <div className="relative w-full h-[85vh]">
  <video
        ref={videoRef}
        className="absolute top-0 left-0 w-full h-full object-cover z-0"
        controls
        src="https://www.w3schools.com/html/mov_bbb.mp4"
        onTimeUpdate={handleTimeUpdate} // Called whenever the time updates
        autoPlay
        loop
      ></video>
  {currentAd && (
    <div
  className="absolute flex items-center justify-center cursor-pointer w-[40vw] h-[40vh] border-4 border-white bg-gray-800 shadow-xl hover:scale-110 transition-transform duration-300"
  style={{ top: adPosition.top, left: adPosition.left }}
  onClick={() => handleAdClick(currentAd)}
>
  <img
    src={currentAd.image_url}
    alt="Ad"
    className="w-full h-full object-cover" // Cover the entire container
    onError={(e) => (e.target.src = "/default-ad.jpg")} // Default image on error
  />
</div>

  )}
</div>
  );
};

export default VideoPlayer;
