import "./ItemList.css"
import { useState } from "react"

export default function ItemList() {
  const [count, setCount] = useState(0)

  return (
    <div className="item-list-container">
      <div className="counter">Results found: {count}</div>
    </div>
  )
}
