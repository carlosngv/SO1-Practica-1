import React, { useEffect, useState } from 'react'
import { Link } from "react-router-dom";
import { variable } from '../helpers/variables';

export const CarTable = () => {

    const url = `${process.env.SERVER_URI}/cars/`;
// const url = `http://localhost:9000/cars/`;

    console.log(process.env)
    const [cars, setCars] = useState({
        data: null,
        isLoading: true,
    });

    const { data, isLoading } = cars;

    useEffect(() => {
        getCars();
    }, []);

    const getCars = async () => {

        let resp = await fetch(url);
        let data = await resp.json();

        setCars({
            data,
            isLoading: false,
        })

    }

    const deleteCar = async ( id ) => {
        await fetch(`${ url }delete/${ id }`);

        getCars();

    }

  return (
      <div className="container">

        {
           isLoading ?
            (
                <div className="alert alert-info text-center mt-5">
                    Loading...
                </div>
            )
            :
            (
                <table className="table table-dark m-5">

                <thead>
                    <tr className="table-dark">
                        <th scope="col">Brand</th>
                        <th scope="col">Model</th>
                        <th scope="col">Series</th>
                        <th scope="col">Plate</th>
                        <th scope="col">Color</th>
                        <th scope="col">Update</th>
                        <th scope="col">Delete</th>
                    </tr>
                </thead>
                <tbody>
                {
                        data.map( car => (
                            <tr className="table-dark" key={ car._id}>
                                <td>{ car.brand }</td>
                                <td>{ car.model }</td>
                                <td>{ car.series }</td>
                                <td>{ car.plate }</td>
                                <td>{ car.color }</td>
                                <td> <Link to={'updateCar/'+car._id} className="btn btn-info m-2">Update</Link></td>
                                <td> <button onClick={ () => deleteCar(car._id) } className="btn btn-danger m-2">Delete</button></td>

                            </tr>
                        ))
                    }

                </tbody>
                </table>
            )

        }


      </div>
  )
}
