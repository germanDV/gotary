import React from 'react'
import { createRoot } from "react-dom/client"
import { createHashRouter, RouterProvider } from "react-router-dom"
import { toastConfig } from "react-simple-toasts"
import Auth from "./pages/Auth"
import Home, { loader as homeLoader } from "./pages/Home"
import Err from "./pages/Err"
import "./style.css"

toastConfig({
  time: 3_000,
  position: "top-right",
  clickClosable: true,
  className: "toast",
})

const router = createHashRouter([
  {
    path: "/",
    element: <Auth />,
    errorElement: <Err />,
  },
  {
    path: "/dashboard",
    element: <Home />,
    errorElement: <Err />,
    loader: homeLoader,
  },
])

const root = createRoot(document.getElementById('root'))
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
