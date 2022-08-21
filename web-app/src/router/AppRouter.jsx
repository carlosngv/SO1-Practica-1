import React from 'react'
import { Routes, Route } from 'react-router-dom'
import { CarTable } from '../cars/components/CarTable'
import { NewCarPage } from '../cars/pages/NewCarPage'
import { UpdateCarPage } from '../cars/pages/UpdateCarPage'
import { Navbar } from '../ui/components/Navbar'

export const AppRouter = () => {
  return (
    <>
        <Navbar />

        <Routes>
            <Route path="newCar" element={ <NewCarPage /> } />
            <Route path="updateCar/:id" element={ <UpdateCarPage /> } />
            <Route path="" element={ <CarTable /> } />
        </Routes>
    </>
  )
}
