import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { LoginComponent } from "./login.component"

describe('LoginComponent', () => {
    it('mounts', () => {
        cy.mount(LoginComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      
      it.only('login without response', () => {
        cy.mount(LoginComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy=login_form]').submit()
      })

      it.only('login with response', () => {
        cy.mount(LoginComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy="login_email_data"]').type('email@gmail.com')
      })
})