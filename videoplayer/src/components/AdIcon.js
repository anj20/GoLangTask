import React from "react";

const AdIcon = ({ ad, position, onClick }) => {
  return (
    <div
      className="absolute cursor-pointer w-20 h-20 flex items-center justify-center bg-white border-2 border-gray-300 rounded-full shadow-lg hover:scale-110 transition-transform duration-300"
      style={{ top: position.top, left: position.left }}
      onClick={onClick}
    >
      <img
        src={ad.image_url}
        alt="Ad"
        className="w-full h-full rounded-full object-cover"
      />
    </div>
  );
};

export default AdIcon;
