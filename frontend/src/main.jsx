import React from 'react'
import { createRoot } from "react-dom/client"
import { createHashRouter, RouterProvider } from "react-router-dom"
import { toastConfig } from "react-simple-toasts"
import Auth, { action as authAction } from "./pages/Auth"
import Home, {
  loader as homeLoader,
  action as homeAction,
} from "./pages/Home"
import Success from "./pages/Success"
import Err from "./pages/Err"
import Signer from "./components/Signer"
import Verifier, { action as verifierAction } from "./components/Verifier"
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
    action: authAction,
  },
  {
    path: "/dashboard",
    element: <Home />,
    errorElement: <Err />,
    loader: homeLoader,
    action: homeAction,
    children: [
      {
        path: "sign",
        element: <Signer />,
      },
      {
        path: "verify",
        element: <Verifier />,
        action: verifierAction,
      },
    ],
  },
  {
    path: "/success",
    element: <Success />,
  },
])

const root = createRoot(document.getElementById('root'))
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
