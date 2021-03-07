/**
 * Configure your Gatsby site with this file.
 *
 * See: https://www.gatsbyjs.com/docs/gatsby-config/
 */

module.exports = {
  /* Your site config here */
  siteMetadata: {
    title: 'Sc4ramouche Blog',
    description:
      'A wee corner where I share thoughts, experience, and review books I read.',
  },
  plugins: [
    'gatsby-plugin-react-helmet',
    'gatsby-plugin-mdx',
    {
      resolve: 'gatsby-source-filesystem',
      options: { path: `${__dirname}/posts` },
      name: 'posts',
    },
  ],
  pathPrefix: '/blog',
};
