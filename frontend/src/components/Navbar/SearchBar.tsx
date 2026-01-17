import { useState } from "react";
import "./SearchBar.css";
import icon from "../../assets/search-interface-symbol.png";

interface Props {
  onSearch: (value: string) => void;
}

export default function SearchBar({ onSearch }: Props) {
  const [query, setQuery] = useState("");

  const handleChange = (value: string) => {
    setQuery(value);
    onSearch(value);
  };

  return (
    <div className="search-bar">
      <img className="icon" src={icon} alt="Logo" />
      <input
        type="text"
        placeholder="Search for games"
        value={query}
        onChange={(e) => handleChange(e.target.value)}
      />
      {query && (
        <button className="clear" onClick={() => handleChange("")}>
          ✖
        </button>
      )}
    </div>
  );
}

