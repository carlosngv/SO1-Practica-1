import { useEffect, useState } from "react"

export const useFetch = ( url, method = 'GET', body = {} ) => {
    const [response, setResponse] = useState({
        isLoading: true,
        data: null,
    });

    const getData = async() => {

        setResponse({
            ...response,
            isLoading: true,
        });

        console.log(url);

        const resp = await fetch(url);
        const data = await resp.json();

        setResponse({
            data,
            isLoading: false,
        });
    }

    const postData = async() => {

        setResponse({
            ...response,
            isLoading: true,
        });

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify( JSON.stringify(body) ),
        };

        let resp = await fetch( url, requestOptions );
        let data = await resp.json();
        setResponse({
            data,
            isLoading: false,
        });
    }

    const putData = async() => {

        setResponse({
            ...response,
            isLoading: true,
        });

        const requestOptions = {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify( body )
        };

        let resp = await fetch( url, requestOptions );
        let data = await resp.json();
        setResponse({
            data,
            isLoading: false,
        });
    }

    const deleteData = async() => {

        setResponse({
            ...response,
            isLoading: true,
        });


        let resp = await fetch( url, { method: "DELETE" } );
        let data = await resp.json();

        setResponse({
            data,
            isLoading: false,
        });

    }

    useEffect(() => {

        switch (method) {
            case 'GET':
                getData();
                break;

            case 'POST':
                postData();
                break;
            case 'PUT':
                putData();
                break;

            case 'DELETE':
                deleteData();
                break;
        }

    }, [ url ]);

    return {
        ...response,
        response
    }

}
