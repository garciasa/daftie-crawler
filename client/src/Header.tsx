import React from 'react';

function Header(): React.ReactElement {
  return (
    <section className="header flex items-center mx-auto bg-houseBlue-light h-12">
        <div className="flex-shrink-0 text-white font-semibold pl-2">
          HouseCrawler <span className="text-xs">v0.1</span>
        </div>
        <div className="mx-auto text-white">Wexford Town</div>
    </section>
  );
}

export default Header;