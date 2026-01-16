import { useState } from "react"
import "./SearchBar.css"
import icon from "../../assets/search-interface-symbol.png"

export default function SearchBar() {
  const [query, setQuery] = useState("")

  return (
    <div className="search-bar">
      <img className="icon" src={icon} alt="Logo" />
      <input
        type="text"
        placeholder="Search for games"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />
      {query && (
        <button className="clear" onClick={() => setQuery("")}>
          ✖
        </button>
      )}
    </div>
  )
}
