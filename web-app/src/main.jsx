import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { CarsApp } from './CarsApp'
import './styles.css'



ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <BrowserRouter>
      <CarsApp />
    </BrowserRouter>
  </React.StrictMode>
)
