import React from 'react'
import { useNavigate } from 'react-router-dom'
import { useForm } from '../../ui/hooks/useForm'

export const NewCarPage = () => {

  // const url = `${import.meta.env.SERVER_URI}/cars/`;
  const url = `${process.env.SERVER_URI}/cars/`;

  const navigate = useNavigate();

  const handleSubmit = async ( e ) => {

    e.preventDefault();

    console.log(form);

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify( form ),
    };

    let resp = await fetch( url, requestOptions );
    let data = await resp.json();
    console.log(data);

    onResetForm();
    navigate(-1);
  }

  const {
    brand,
    plate,
    model,
    color,
    series,
    onInputChange,
    onResetForm,
    form
  } = useForm({
    brand: '',
    plate: '',
    model: '',
    color: '',
    series: '',
  })

  return (
    <div className="container pt-3">
      <div className="row">
        <div className="col-4"></div>
        <div className="col-4">
          <h1>New Car</h1>
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
                  Add
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
