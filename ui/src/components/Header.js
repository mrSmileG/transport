import React from 'react';
import { Menu } from 'semantic-ui-react';
import { Link } from '../routes';

export default () => {
    return (
        <Menu style={{ marginTop: '10px' }}>
            <Link route="/">
                <a href="." className="item">Transportation Problem Solver</a>
            </Link>

            <Menu.Menu position="right">
                <Link route="/">
                    <a href="." className="item">+</a>
                </Link>
            </Menu.Menu>
        </Menu>
    );
};