import React from 'react'
import "./card-bar.css";

function CardBar() {
    return (
    <div class="cardbar">
        <div class="left">
            <div class="usericon">
            </div>

            <ul class="userlink">
            <li><a href="https://letmegooglethat.com/?q=why+doesn%27t+chagy+pherome">@jasontduong</a></li>
            </ul>
        </div>

        <div class="right">
          <button class="settingscardbar">
            settings
          </button>
        
          <div class="socialicon">

          </div>
        </div>

    </div>
    );
  }

  export default CardBar;