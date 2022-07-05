import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";

const client = new ApolloClient({
    cache: new InMemoryCache(),
    defaultOptions: {
        watchQuery: {
            fetchPolicy: "no-cache",
            errorPolicy: "ignore",
        },
        query: {
            fetchPolicy: "no-cache",
            errorPolicy: "all",
        },
    },
});

ReactDOM.createRoot(document.getElementById('root')!).render(
    <ApolloProvider client={client}>
    <App />
  </ApolloProvider>
)
