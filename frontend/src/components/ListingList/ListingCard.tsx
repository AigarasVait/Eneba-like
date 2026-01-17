import React from "react";
import "./ListingCard.css"
import type { ListingCardDTO } from "../../types/ListingCardDTO";

interface Props {
  listing: ListingCardDTO;
}

const ListingCard: React.FC<Props> = ({ listing }) => {
  return (
    <div className="card">
      <img src={listing.gameImagePath} alt={listing.name} className="cover" />
      <div className="info">
        <h3>{listing.name}</h3>
        <p>{listing.region}</p>
        <p className="price">
          <span className="original">€{listing.basePrice.toFixed(2)}</span>
          <span className="discount">-{listing.discountPercent}%</span>
          <span className="final">€{listing.price.toFixed(2)}</span>
        </p>
        <p className="cashback">Cashback: €{listing.cashback.toFixed(2)}</p>
        <p className="likes">❤️ {listing.favoritedCount}</p>
      </div>
    </div>
  );
};

export default ListingCard;
