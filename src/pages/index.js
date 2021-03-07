import React from 'react';
import '@fontsource/ubuntu';

import Page from '../components/Page';
import SEO from '../components/SEO';
import Button from '../components/Button';
import { graphql, Link } from 'gatsby';

export default function Home({ data }) {
  const posts = data.allMdx.nodes;

  return (
    <Page>
      <SEO title={'Home'} />
      <div
        style={{
          color: `#333`,
          fontSize: `56px`,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <h1>Hey there 👋</h1>
        <p>I'm Vlad. Here I'll share my thoughts and ideas. Soon!</p>
        {posts.map(({ frontmatter: { title }, slug }) => (
          <Link to={slug} key={slug}>
            <h4>{title}</h4>
          </Link>
        ))}
      </div>
    </Page>
  );
}

export const postsQuery = graphql`
  {
    allMdx(sort: { fields: [frontmatter___title], order: ASC }) {
      nodes {
        slug
        frontmatter {
          title
        }
      }
    }
  }
`;
