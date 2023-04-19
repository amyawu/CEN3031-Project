import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AccountComponent } from "./account.component"

describe('Account Page without login', () => {
    it('mounts', () => {
        cy.mount(AccountComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      
      it.only('Account page with login', () => { // need to fix
        cy.mount(AccountComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy=profile_form]').submit()
      })
})