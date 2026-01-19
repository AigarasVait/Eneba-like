import { useState } from "react";
import Navbar from "../components/Navbar/Navbar";
import ListingGrid from "../components/ListingGrid/ListingGrid";
import "./HomePage.css";

export default function HomePage() {
  const [search, setSearch] = useState("");

  return (
    <div>
      <Navbar onSearch={setSearch} />
      <ListingGrid search={search} />
    </div>
  );
}