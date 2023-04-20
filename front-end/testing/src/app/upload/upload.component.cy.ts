import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { UploadComponent } from "./upload.component"

describe('UploadComponent', () => {
    it('mounts', () => {
        cy.mount(UploadComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      
      it.only('check if you can upload', () => {
        cy.mount(UploadComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy=upload_button]').click()
      })
})