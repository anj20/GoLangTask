import React from "react";
import "./AdIcon.css";

const AdIcon = ({ ad, position, onClick }) => {
  return (
    <div
      className="ad-icon"
      style={{ top: position.top, left: position.left }}
      onClick={onClick}
    >
      <img src={ad.image_url} alt="Ad" />
    </div>
  );
};

export default AdIcon;
