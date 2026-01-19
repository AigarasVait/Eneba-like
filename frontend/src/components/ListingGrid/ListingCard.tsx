import React from "react";
import "./ListingCard.css"
import type { ListingCardDTO } from "../../types/ListingCardDTO";
import { API_BASE_URL } from "../../constants/api";
import heart from "../../assets/heart.png"

interface Props {
  listing: ListingCardDTO;
}

const ListingCard: React.FC<Props> = ({ listing }) => {
  return (
    <div className="card">
      <div className="cover-wrapper">
        <img 
        src={listing.gameImagePath ? `${API_BASE_URL}/images/${listing.gameImagePath}` : ''} 
        alt={listing.name} 
        className="cover" 
        />
        <div className="store-overlay">
          <img 
          src={listing.gameStoreImagePath ? `${API_BASE_URL}/images/${listing.gameStoreImagePath}` : ''} 
          alt={listing.gameStoreName} 
          className="store-cover" 
          />
        ️  <span>{listing.gameStoreName}</span>
        </div>
        <div className="cashback-overlay">
          <span>Cashback</span>
        </div>
      </div>
      <div className="info">
        <div className="top">
          <p className="title">{listing.name}</p>
          <p className="region">{listing.region}</p>
        </div>
        
        <div className="bottom">
          <p className="price">
            <span className="original">From</span>
            {Number(listing.basePrice) !== Number(listing.price) && (
              <>
                <span className="original-number">
                  €{listing.basePrice?.toFixed(2)}
                </span>

                <span className="discount">
                  -{listing.discountPercent?.toFixed(0) ?? 0}%
                </span>
              </>
            )}

          </p>
          <p className="final">€{listing.price?.toFixed(2) ?? '0.00'}</p>
          <p className="discount">Cashback: €{listing.cashback?.toFixed(2) ?? '0.00'}</p>
          <p className="likes"> 
            <img className="like-icon" src={heart} alt="Likes" />
            <span>{listing.favoritedCount ?? 0}</span>
          </p>
        </div>
      </div>
    </div>
  );
};

export default ListingCard;
