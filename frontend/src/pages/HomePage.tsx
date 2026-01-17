import { useState } from "react";
import Navbar from "../components/Navbar/Navbar";
import ListingGrid from "../components/ListingGrid/ListingGrid";

export default function HomePage() {
  const [search, setSearch] = useState("");

  return (
    <>
      <Navbar onSearch={setSearch} />
      <ListingGrid search={search} />
    </>
  );
}
