import { HttpClientModule } from "@angular/common/http"
import { LoginComponent } from "./login.component"

describe('LoginComponent', () => {
    it('mounts', () => {
        cy.mount(LoginComponent, {
          imports: [HttpClientModule]
        })
      })
})