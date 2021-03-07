import React from 'react';
import '@fontsource/ubuntu';

import Page from '../components/Page';
import Button from '../components/Button';

export default function Home() {
  return (
    <Page>
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
      </div>
    </Page>
  );
}
