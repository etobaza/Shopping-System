import React, { useEffect, useState } from 'react';
import Navbar from './Navbar';
import { fetchItems } from '../services/item';
import { useParams } from 'react-router-dom';

const CategoryList = () => {
    const { category } = useParams();
    const [items, setItems] = useState([]);

    useEffect(() => {
        console.log('Fetching items for category:', category);

        fetchItems().then(items => {
            const filteredItems = items.filter(item => item.category === category);
            console.log('Filtered items:', filteredItems);
            setItems(filteredItems);
        });
    }, [category]);

    console.log('Rendering items:', items);

    return (
        <div>
            <Navbar />
            <h2>Items in {category}</h2>
            <ul>
                {items.map(item => (
                    <li key={item.id}>{item.name}</li>
                ))}
            </ul>
        </div>
    );
};

export default CategoryList;
