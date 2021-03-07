import React from 'react';
import { graphql } from 'gatsby';
import { MDXRenderer } from 'gatsby-plugin-mdx';

import Page from '../components/Page';

const BlogPostPage = ({ data }) => {
  const {
    frontmatter: { title, date },
    body,
  } = data.mdx;

  return (
    <Page>
      <article>
        <h1>{title}</h1>
        <span>{date}</span>
        <MDXRenderer>{body}</MDXRenderer>
      </article>
    </Page>
  );
};

export const query = graphql`
  query BlogPostPage($id: String) {
    mdx(id: { eq: $id }) {
      frontmatter {
        title
        date
      }
      body
    }
  }
`;

export default BlogPostPage;
