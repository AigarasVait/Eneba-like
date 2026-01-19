import React, { useEffect, useState } from "react";
import ListingCard from "./ListingCard";
import "./ListingGrid.css";
import type { ListingCardDTO } from "../../types/ListingCardDTO";
import { getListings, searchListings } from "../../services/ListingService.tsx";


interface Props {
  search?: string;
}

const ListingGrid: React.FC<Props> = ({ search }) => {
  const [listings, setListings] = useState<ListingCardDTO[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const load = async () => {
      setLoading(true);
      setError(null);

      try {
        const data =
          search && search.trim().length > 0
            ? await searchListings(search)
            : await getListings();

        setListings(Array.isArray(data) ? data : []);
      } catch (err) {
        console.error("Failed to load listings:", err);
        setError("Failed to load listings. Please try again.");
      } finally {
        setLoading(false);
      }
    };

    load();
  }, [search]);

  if (loading) {
    return <div>Loading…</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  if (
    search &&
    search.trim().length > 0 &&
    listings.length === 0 &&
    !loading
  ) {
    return (
      <div className="no-results">
        No results found for "<strong>{search}</strong>"
      </div>
    );
  }

  return (
    <div className="grid-container">
      <div className="results-count">
        {listings.length === 1 ? "Result" : "Results"} found: {listings.length}
      </div>

      <div className="grid">
        {listings.map((listing) => (
          <ListingCard key={listing.id} listing={listing} />
        ))}
      </div>
    </div>
  );
};

export default ListingGrid;
