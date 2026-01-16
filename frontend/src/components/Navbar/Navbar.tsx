import "./Navbar.css"
import SearchBar from "./SearchBar"
import logo from "../../assets/logo.png"
import heart from "../../assets/heart.png"
import cart from "../../assets/shopping-cart.png"
import user from "../../assets/user.svg"
import flag from "../../assets/lithuania.svg"

export default function Navbar() {
  return (
    <nav className="navbar">
      <div className="container">
        <img className="logo" src={logo} alt="Logo" />
        <SearchBar />
        <div className="container">
          <img className="flag" src={flag} alt="flag" />
          <span>English EU | EUR</span>
        </div>
      </div>

      <div className="container">
        <img className="icon" src={heart} alt="Favorites" />
        <img className="icon" src={cart} alt="Cart" />
        <img className="user-icon" src={user} alt="User" />
      </div>
    </nav>
  )
}
