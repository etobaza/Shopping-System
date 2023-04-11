import React, { useEffect, useState } from 'react';
import Navbar from './Navbar';
import { fetchItems } from '../services/item';
import { useParams, useNavigate, useLocation } from 'react-router-dom';

const CategoryList = () => {
    const { category } = useParams();
    const [items, setItems] = useState([]);
    const navigate = useNavigate();
    const location = useLocation();
    const userType = localStorage.getItem('userType');
    const searchParams = new URLSearchParams(location.search);
    const searchQuery = searchParams.get('q') || '';
    const priceFilter = searchParams.get('price') || '';
    const ratingFilter = searchParams.get('rating') || '';

    useEffect(() => {
        console.log('Fetching items for category:', category);

        fetchItems().then(items => {
            let filteredItems = items.filter(item => item.category === category);
            if (searchQuery !== '') {
                filteredItems = filteredItems.filter(item => item.name.toLowerCase().includes(searchQuery.toLowerCase()));
            }

            if (priceFilter !== '') {
                filteredItems = filteredItems.filter(item => item.price <= parseInt(priceFilter));
            }

            if (ratingFilter !== '') {
                filteredItems = filteredItems.filter(item => item.rating >= parseInt(ratingFilter));
            }

            console.log('Filtered items:', filteredItems);
            setItems(filteredItems);
        });
    }, [category, searchQuery, priceFilter, ratingFilter]);

    console.log('Rendering items:', items);

    const handleBuyClick = (item) => {
        let currentCart = JSON.parse(localStorage.getItem('cartItems')) || [];
        currentCart.push(item);
        localStorage.setItem('cartItems', JSON.stringify(currentCart));
    };

    const handleEditClick = (itemId) => {
        navigate(`/seller/items/${itemId}`);
    };

    const renderItem = (item) => {
        return (
            <div key={item.id}>
                <h3>{item.name}</h3>
                <p>Price: {item.price}</p>
                <p>Rating: {item.rating}</p>
                {userType === 'customer' && (
                    <button onClick={() => handleBuyClick(item)}>Buy</button>
                )}
                {userType === 'seller' && (
                    <div>
                        <button onClick={() => handleEditClick(item.id)}>Edit</button>
                    </div>
                )}
            </div>
        );
    };

    return (
        <div>
            <Navbar />
            <h2>Items in {category}</h2>
            <form>
                <label>
                    Search:
                    <input type="text" name="q" defaultValue={searchQuery} />
                </label>
                <label>
                    Max Price:
                    <input type="number" name="price" defaultValue={priceFilter} />
                </label>
                <label>
                    Min Rating:
                    <input type="number" name="rating" defaultValue={ratingFilter} />
                </label>
                <button type="submit">Filter</button>
            </form>
            {items.length > 0 ? (
                <div>{items.map(renderItem)}</div>
            ) : (
                <p>No items found.</p>
            )}
        </div>
    );
};

export default CategoryList;
