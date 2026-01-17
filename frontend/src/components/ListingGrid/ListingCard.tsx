import React from "react";
import "./ListingCard.css"
import type { ListingCardDTO } from "../../types/ListingCardDTO";
import { API_BASE_URL } from "../../constants/api";

interface Props {
  listing: ListingCardDTO;
}

const ListingCard: React.FC<Props> = ({ listing }) => {
  return (
    <div className="card">
      <img 
        src={listing.gameImagePath ? `${API_BASE_URL}/images/${listing.gameImagePath}` : ''} 
        alt={listing.name} 
        className="cover" 
      />
      <div className="info">
        <h3>{listing.name}</h3>
        <p>{listing.region}</p>
        <p className="price">
          <span className="original">€{listing.basePrice?.toFixed(2) ?? '0.00'}</span>
          <span className="discount">-{listing.discountPercent ?? 0}%</span>
          <span className="final">€{listing.price?.toFixed(2) ?? '0.00'}</span>
        </p>
        <p className="cashback">Cashback: €{listing.cashback?.toFixed(2) ?? '0.00'}</p>
        <p className="likes">❤️ {listing.favoritedCount ?? 0}</p>
      </div>
    </div>
  );
};

export default ListingCard;
