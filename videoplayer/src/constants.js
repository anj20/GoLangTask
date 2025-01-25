// src/constants.js
export const API_BASE_URL = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080";
export const AD_FETCH_INTERVAL = parseInt(process.env.REACT_APP_AD_FETCH_INTERVAL, 10) || 10000;

export const VIDEO_PLAYER_STYLE = {
  width: "100%",
  height: "auto",
  backgroundColor: "#000",
};

export const AD_ICON_STYLE = {
  position: "absolute",
  cursor: "pointer",
  zIndex: 10,
};

export const AD_POSITIONS = [
  { top: "10%", left: "10%" },
  { top: "20%", left: "60%" },
  { top: "60%", left: "20%" },
  { top: "60%", left: "60%" },
];
