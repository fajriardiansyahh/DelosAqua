

$(function () {
    $('button.btn-farms').each(function () {
        $(this).on('click', function () {
            let id = $(this).attr('farms-data')
            window.location = '/Farms?id='+ id
        })
    })
})