import React, { useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useForm } from '../../ui/hooks/useForm'

export const UpdateCarPage = () => {

   const { id } = useParams();

  //  const url = `http://localhost:9001/cars/`;
   const url = `${process.env.SERVER_URI}/cars/`;
  const {
      brand,
      plate,
      model,
      color,
      series,
      onInputChange,
      onResetForm,
      form,
      setForm
    } = useForm({
        brand: '',
        plate: '',
        model: '',
        color: '',
        series: '',
    });

    const fetchCar = async () => {
        let resp = await fetch(url + 'car/' + id);
        let data = await resp.json();
        setForm({
            ...form,
            brand: data.brand,
            model: data.model,
            color: data.color,
            plate: data.plate,
            series: data.series,
        })

      }

  useEffect(() => {
      fetchCar();
  }, []);

  const navigate = useNavigate();

  const handleSubmit = async ( e ) => {

    e.preventDefault();

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify( form ),
    };



    let resp = await fetch( url + "update/" + id, requestOptions );
    let data = await resp.json();


    onResetForm();
    navigate(-1);
  }


  return (
    <div className="container pt-3">
      <div className="row">
        <div className="col-4"></div>
        <div className="col-4">
          <h1>Update Car</h1>
          <hr />

          <form onSubmit={ handleSubmit }>
            <div className="form-group">
              <input
                type="text"
                placeholder="Brand"
                className="form-control mt-2"
                name="brand"
                value={ brand }
                autoComplete="off"
                onChange={ onInputChange }
              />

              <input
                type="text"
                placeholder="Model"
                className="form-control mt-2"
                name="model"
                value={ model }
                autoComplete="off"
                onChange={ onInputChange }
              />

              <input
                type="text"
                placeholder="Series"
                className="form-control mt-2"
                name="series"
                value={ series }
                autoComplete="off"
                onChange={ onInputChange }
              />

              <input
                type="text"
                placeholder="Plate"
                className="form-control mt-2"
                name="plate"
                value={ plate }
                autoComplete="off"
                onChange={ onInputChange }
              />

              <input
                type="text"
                placeholder="Color"
                className="form-control mt-2"
                name="color"
                value={ color }
                autoComplete="off"
                onChange={ onInputChange }
              />
              <div className="d-grid gap-2">
                <button className="btn btn-outline-primary mt-2">
                  Update
                </button>
              </div>

            </div>

        </form>

        </div>
        <div className="col-4"></div>
      </div>

    </div>
  )
}
