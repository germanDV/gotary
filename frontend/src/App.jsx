import { createHashRouter, RouterProvider, Link } from "react-router-dom"
import "./App.css"

const router = createHashRouter([
  {
    path: "/",
    element: <Auth />,
  },
  {
    path: "/dashboard",
    element: <Home />,
  },
])

function App() {
  return <RouterProvider router={router} />
}

export default App
