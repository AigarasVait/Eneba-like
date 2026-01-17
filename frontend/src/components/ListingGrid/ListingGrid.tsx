import React, { useEffect, useState } from "react";
import ListingCard from "./ListingCard";
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
        const data = search && search.trim().length > 0
          ? await searchListings(search)
          : await getListings();

        setListings(data);
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

  return (
    <div className="grid">
      {listings.map(listing => (
        <ListingCard key={listing.id} listing={listing} />
      ))}
    </div>
  );
};

export default ListingGrid;