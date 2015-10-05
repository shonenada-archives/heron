$('#signin-form').ajaxForm
  dataType: 'json'
  type: 'POST'
  success: (resp) ->
    if resp.success
      location.href = '/';
    else
      $('#tips').html(resp.message)

$('#signup-form').ajaxForm
  dataType: 'json'
  type: 'POST'
  success: (resp) ->
    if resp.success
      location.href = '/account/signin'
    else
      $('#tips').html(resp.message)
