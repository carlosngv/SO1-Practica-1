import React, { useEffect } from 'react'
import { Link } from "react-router-dom";
import { useForm } from '../hooks/useForm';

export const Navbar = () => {

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark p-3">
        <Link
            className="navbar-brand"
            to="/">
            Cars Application
        </Link>

        <Link
          className="btn btn-success btn-lg"
          to="/newCar">
            New
        </Link>




        <div className="navbar-collapse collapse w-100 order-3 dual-collapse2 d-flex justify-content-end">
          <span className="nav-item text-light">Carlos Ng - 20180434</span>
        </div>
    </nav>
  )
}
