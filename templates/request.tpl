package {{.BasePackage}}.model.request;

import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.media.Schema.RequiredMode;
import java.math.BigDecimal;
import java.util.List;
import javax.validation.constraints.Digits;
import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.NotNull;
import javax.validation.constraints.Pattern;
import javax.validation.constraints.Positive;
import lombok.Getter;
import lombok.Setter;

/**
 * {{.EntityDescription}} request
 *
 * @author: {{.Author}}
 */
@Getter
@Setter
@Schema(description = "{{.EntityDescription}}")
public class {{.EntityName}}Request {

{{ range $index, $field := .EntityFields }}
  /**
   * {{$field.Description}}
   */
  {{if $field.NotNull}}@Schema(description = "{{$field.Description}}", requiredMode = RequiredMode.REQUIRED){{else }}@Schema(description = "{{$field.Description}}"){{end}}
  {{if and (eq $field.Type "String") ($field.NotNull)}}@NotEmpty(message = "{{$field.Description}} 不能为空"){{else if and (ne $field.Type "String") ($field.NotNull)}}@NotNull(message = "{{$field.Description}} 不能为空"){{end}}
  private {{$field.Type}} {{$field.Name}};
{{ end }}

}