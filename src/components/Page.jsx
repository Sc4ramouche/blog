import React from 'react';
import { Link } from 'gatsby';

const Page = ({ children }) => (
  <>
    <header>
      <nav>
        <Link to="/">Home</Link>
      </nav>
    </header>
    <main>{children}</main>
    <footer>All rights reserved</footer>
  </>
);

export default Page;
